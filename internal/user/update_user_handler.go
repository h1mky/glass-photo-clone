package user

import (
	"database/sql"
	"log"
	"net/http"

	"glass-photo/internal/common"
)

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	targetUserID, err := common.GetParams(r)
	if err != nil {
		log.Printf("ERROR: updateUserHandler - Invalid user ID in URL: %v", err)
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	requesterUserIDVal := ctx.Value("user_id")
	if requesterUserIDVal == nil {
		log.Println("ERROR: updateUserHandler - Authenticated user ID not found in context (middleware issue).")
		http.Error(w, "Authentication required", http.StatusUnauthorized)
		return
	}
	requesterUserID, ok := requesterUserIDVal.(int)
	if !ok {
		log.Printf("ERROR: updateUserHandler - Authenticated user ID has wrong type: %T", requesterUserIDVal)
		http.Error(w, "Internal server error: Authentication data corrupted", http.StatusInternalServerError)
		return
	}

	if requesterUserID != targetUserID {
		log.Printf("WARN: User %d attempted to update profile of user %d without permission.", requesterUserID, targetUserID)
		http.Error(w, "Forbidden: You are not authorized to update this profile.", http.StatusForbidden)
		return
	}

	var patchRequest UserPatchRequest
	if err := common.JsonParse(w, r, &patchRequest); err != nil {

		log.Printf("ERROR: updateUserHandler - Failed to parse JSON: %v", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	updatedProfile, err := updateUserAvatarAndDescriptionPostgres(ctx, targetUserID, patchRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("INFO: User with ID %d not found for update.", targetUserID)
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err.Error() == "no fields provided for update (user_img or description)" {
			http.Error(w, "No valid fields (user_img or description) provided for update.", http.StatusBadRequest)
			return
		}
		log.Printf("ERROR: updateUserHandler - Failed to update user profile %d: %v", targetUserID, err)
		http.Error(w, "Failed to update user profile due to server error.", http.StatusInternalServerError)
		return
	}

	common.JsonEncode(w, updatedProfile)
	log.Printf("INFO: User %d profile updated successfully.", targetUserID)
}
