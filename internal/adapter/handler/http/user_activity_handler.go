// internal/adapter/handler/http/user_activity_handler.go
package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yvanyang/language-learning-player-api/internal/adapter/handler/http/dto"
	"github.com/yvanyang/language-learning-player-api/internal/adapter/handler/http/middleware"
	"github.com/yvanyang/language-learning-player-api/internal/domain"
	"github.com/yvanyang/language-learning-player-api/internal/port"
	"github.com/yvanyang/language-learning-player-api/pkg/httputil"
	"github.com/yvanyang/language-learning-player-api/pkg/pagination"
	"github.com/yvanyang/language-learning-player-api/pkg/validation"
)

// UserActivityHandler handles HTTP requests related to user progress and bookmarks.
type UserActivityHandler struct {
	activityUseCase port.UserActivityUseCase // 使用port包中定义的接口
	validator       *validation.Validator
}

// NewUserActivityHandler creates a new UserActivityHandler.
func NewUserActivityHandler(uc port.UserActivityUseCase, v *validation.Validator) *UserActivityHandler {
	return &UserActivityHandler{
		activityUseCase: uc,
		validator:       v,
	}
}

// --- Progress Handlers ---

// RecordProgress handles POST /api/v1/users/me/progress
// @Summary Record playback progress
// @Description Records or updates the playback progress for a specific audio track for the authenticated user.
// @ID record-playback-progress
// @Tags User Activity
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param progress body dto.RecordProgressRequestDTO true "Playback progress details"
// @Success 204 "Progress recorded successfully"
// @Failure 400 {object} httputil.ErrorResponseDTO "Invalid Input / Track ID Format"
// @Failure 401 {object} httputil.ErrorResponseDTO "Unauthorized"
// @Failure 404 {object} httputil.ErrorResponseDTO "Track Not Found"
// @Failure 500 {object} httputil.ErrorResponseDTO "Internal Server Error"
// @Router /users/me/progress [post]
func (h *UserActivityHandler) RecordProgress(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		httputil.RespondError(w, r, domain.ErrUnauthenticated)
		return
	}

	var req dto.RecordProgressRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: invalid request body", domain.ErrInvalidArgument))
		return
	}
	defer r.Body.Close()

	if err := h.validator.ValidateStruct(req); err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: %v", domain.ErrInvalidArgument, err))
		return
	}

	trackID, err := domain.TrackIDFromString(req.TrackID)
	if err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: invalid track ID format", domain.ErrInvalidArgument))
		return
	}

	// Convert milliseconds (int64) to duration
	progressDuration := time.Duration(req.ProgressMs) * time.Millisecond

	err = h.activityUseCase.RecordPlaybackProgress(r.Context(), userID, trackID, progressDuration)
	if err != nil {
		httputil.RespondError(w, r, err) // Handles NotFound (track), internal errors
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content is suitable for successful upsert
}

// GetProgress handles GET /api/v1/users/me/progress/{trackId}
// @Summary Get playback progress for a track
// @Description Retrieves the playback progress for a specific audio track for the authenticated user.
// @ID get-playback-progress
// @Tags User Activity
// @Produce json
// @Security BearerAuth
// @Param trackId path string true "Audio Track UUID" Format(uuid)
// @Success 200 {object} dto.PlaybackProgressResponseDTO "Playback progress found"
// @Failure 400 {object} httputil.ErrorResponseDTO "Invalid Track ID Format"
// @Failure 401 {object} httputil.ErrorResponseDTO "Unauthorized"
// @Failure 404 {object} httputil.ErrorResponseDTO "Progress Not Found (or Track Not Found)"
// @Failure 500 {object} httputil.ErrorResponseDTO "Internal Server Error"
// @Router /users/me/progress/{trackId} [get]
func (h *UserActivityHandler) GetProgress(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		httputil.RespondError(w, r, domain.ErrUnauthenticated)
		return
	}

	trackIDStr := chi.URLParam(r, "trackId")
	trackID, err := domain.TrackIDFromString(trackIDStr)
	if err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: invalid track ID format", domain.ErrInvalidArgument))
		return
	}

	progress, err := h.activityUseCase.GetPlaybackProgress(r.Context(), userID, trackID)
	if err != nil {
		// Handles ErrNotFound correctly via RespondError
		httputil.RespondError(w, r, err)
		return
	}

	resp := dto.MapDomainProgressToResponseDTO(progress) // Mapper converts duration to ms
	httputil.RespondJSON(w, r, http.StatusOK, resp)
}

// ListProgress handles GET /api/v1/users/me/progress
// @Summary List user's playback progress
// @Description Retrieves a paginated list of playback progress records for the authenticated user.
// @ID list-playback-progress
// @Tags User Activity
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Pagination limit" default(50) minimum(1) maximum(100)
// @Param offset query int false "Pagination offset" default(0) minimum(0)
// @Success 200 {object} dto.PaginatedResponseDTO{data=[]dto.PlaybackProgressResponseDTO} "Paginated list of playback progress"
// @Failure 401 {object} httputil.ErrorResponseDTO "Unauthorized"
// @Failure 500 {object} httputil.ErrorResponseDTO "Internal Server Error"
// @Router /users/me/progress [get]
func (h *UserActivityHandler) ListProgress(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		httputil.RespondError(w, r, domain.ErrUnauthenticated)
		return
	}

	q := r.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))   // Use 0 if parsing fails
	offset, _ := strconv.Atoi(q.Get("offset")) // Use 0 if parsing fails

	// Create pagination parameters and apply defaults/constraints
	pageParams := pagination.NewPageFromOffset(limit, offset)

	// Create use case parameters struct
	ucParams := port.ListProgressInput{
		UserID: userID,
		Page:   pageParams,
	}

	// Call use case with the params struct
	progressList, total, actualPageInfo, err := h.activityUseCase.ListUserProgress(r.Context(), ucParams)
	if err != nil {
		httputil.RespondError(w, r, err)
		return
	}

	respData := make([]dto.PlaybackProgressResponseDTO, len(progressList))
	for i, p := range progressList {
		respData[i] = dto.MapDomainProgressToResponseDTO(p) // Mapper converts duration to ms
	}

	// Use the returned actualPageInfo for the response
	paginatedResult := pagination.NewPaginatedResponse(respData, total, actualPageInfo)

	// Use the generic PaginatedResponseDTO from the DTO package
	resp := dto.PaginatedResponseDTO{
		Data:       paginatedResult.Data,
		Total:      paginatedResult.Total,
		Limit:      paginatedResult.Limit,
		Offset:     paginatedResult.Offset,
		Page:       paginatedResult.Page,
		TotalPages: paginatedResult.TotalPages,
	}

	httputil.RespondJSON(w, r, http.StatusOK, resp)
}

// --- Bookmark Handlers ---

// CreateBookmark handles POST /api/v1/users/me/bookmarks
// @Summary Create a bookmark
// @Description Creates a new bookmark at a specific timestamp in an audio track for the authenticated user.
// @ID create-bookmark
// @Tags User Activity
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param bookmark body dto.CreateBookmarkRequestDTO true "Bookmark details"
// @Success 201 {object} dto.BookmarkResponseDTO "Bookmark created successfully"
// @Failure 400 {object} httputil.ErrorResponseDTO "Invalid Input / Track ID Format"
// @Failure 401 {object} httputil.ErrorResponseDTO "Unauthorized"
// @Failure 404 {object} httputil.ErrorResponseDTO "Track Not Found"
// @Failure 500 {object} httputil.ErrorResponseDTO "Internal Server Error"
// @Router /users/me/bookmarks [post]
func (h *UserActivityHandler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		httputil.RespondError(w, r, domain.ErrUnauthenticated)
		return
	}

	var req dto.CreateBookmarkRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: invalid request body", domain.ErrInvalidArgument))
		return
	}
	defer r.Body.Close()

	if err := h.validator.ValidateStruct(req); err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: %v", domain.ErrInvalidArgument, err))
		return
	}

	trackID, err := domain.TrackIDFromString(req.TrackID)
	if err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: invalid track ID format", domain.ErrInvalidArgument))
		return
	}

	// Convert milliseconds (int64) to duration
	timestampDuration := time.Duration(req.TimestampMs) * time.Millisecond

	bookmark, err := h.activityUseCase.CreateBookmark(r.Context(), userID, trackID, timestampDuration, req.Note)
	if err != nil {
		httputil.RespondError(w, r, err) // Handles NotFound (track), internal errors
		return
	}

	resp := dto.MapDomainBookmarkToResponseDTO(bookmark) // Mapper converts duration to ms
	httputil.RespondJSON(w, r, http.StatusCreated, resp) // 201 Created
}

// ListBookmarks handles GET /api/v1/users/me/bookmarks
// @Summary List user's bookmarks
// @Description Retrieves a paginated list of bookmarks for the authenticated user, optionally filtered by track ID.
// @ID list-bookmarks
// @Tags User Activity
// @Produce json
// @Security BearerAuth
// @Param trackId query string false "Filter by Audio Track UUID" Format(uuid)
// @Param limit query int false "Pagination limit" default(50) minimum(1) maximum(100)
// @Param offset query int false "Pagination offset" default(0) minimum(0)
// @Success 200 {object} dto.PaginatedResponseDTO{data=[]dto.BookmarkResponseDTO} "Paginated list of bookmarks"
// @Failure 400 {object} httputil.ErrorResponseDTO "Invalid Track ID Format (if provided)"
// @Failure 401 {object} httputil.ErrorResponseDTO "Unauthorized"
// @Failure 500 {object} httputil.ErrorResponseDTO "Internal Server Error"
// @Router /users/me/bookmarks [get]
func (h *UserActivityHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		httputil.RespondError(w, r, domain.ErrUnauthenticated)
		return
	}

	q := r.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))

	// Create pagination parameters
	pageParams := pagination.NewPageFromOffset(limit, offset)

	// Check for optional trackId filter
	var trackIDFilter *domain.TrackID
	if trackIDStr := q.Get("trackId"); trackIDStr != "" {
		tid, err := domain.TrackIDFromString(trackIDStr)
		if err != nil {
			httputil.RespondError(w, r, fmt.Errorf("%w: invalid trackId query parameter format", domain.ErrInvalidArgument))
			return
		}
		trackIDFilter = &tid
	}

	// Create use case parameters struct
	ucParams := port.ListBookmarksInput{
		UserID:        userID,
		TrackIDFilter: trackIDFilter,
		Page:          pageParams,
	}

	// Call use case with the params struct
	bookmarks, total, actualPageInfo, err := h.activityUseCase.ListBookmarks(r.Context(), ucParams)
	if err != nil {
		httputil.RespondError(w, r, err)
		return
	}

	respData := make([]dto.BookmarkResponseDTO, len(bookmarks))
	for i, b := range bookmarks {
		respData[i] = dto.MapDomainBookmarkToResponseDTO(b) // Mapper converts duration to ms
	}

	// Use the returned actualPageInfo for the response
	paginatedResult := pagination.NewPaginatedResponse(respData, total, actualPageInfo)

	// Use the generic PaginatedResponseDTO from the DTO package
	resp := dto.PaginatedResponseDTO{
		Data:       paginatedResult.Data,
		Total:      paginatedResult.Total,
		Limit:      paginatedResult.Limit,
		Offset:     paginatedResult.Offset,
		Page:       paginatedResult.Page,
		TotalPages: paginatedResult.TotalPages,
	}

	httputil.RespondJSON(w, r, http.StatusOK, resp)
}

// DeleteBookmark handles DELETE /api/v1/users/me/bookmarks/{bookmarkId}
// @Summary Delete a bookmark
// @Description Deletes a specific bookmark owned by the current user.
// @ID delete-bookmark
// @Tags User Activity
// @Produce json
// @Security BearerAuth
// @Param bookmarkId path string true "Bookmark UUID" Format(uuid)
// @Success 204 "Bookmark deleted successfully"
// @Failure 401 {object} httputil.ErrorResponseDTO "Unauthorized"
// @Failure 403 {object} httputil.ErrorResponseDTO "Forbidden (Not Owner)"
// @Failure 404 {object} httputil.ErrorResponseDTO "Bookmark Not Found"
// @Failure 500 {object} httputil.ErrorResponseDTO "Internal Server Error"
// @Router /users/me/bookmarks/{bookmarkId} [delete]
func (h *UserActivityHandler) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		httputil.RespondError(w, r, domain.ErrUnauthenticated)
		return
	}

	bookmarkIDStr := chi.URLParam(r, "bookmarkId")
	bookmarkID, err := domain.BookmarkIDFromString(bookmarkIDStr)
	if err != nil {
		httputil.RespondError(w, r, fmt.Errorf("%w: invalid bookmark ID format", domain.ErrInvalidArgument))
		return
	}

	err = h.activityUseCase.DeleteBookmark(r.Context(), userID, bookmarkID)
	if err != nil {
		httputil.RespondError(w, r, err) // Handles NotFound, PermissionDenied
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
