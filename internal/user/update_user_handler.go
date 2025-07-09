package user

import (
	"database/sql"
	"log"
	"net/http"

	"glass-photo/internal/common"
)

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idVal := r.Context().Value("user_id")
	requesterUserID, ok := idVal.(int)
	if !ok {
		log.Println("ERROR: Authenticated user ID not found or invalid type.")
		http.Error(w, "Authentication required", http.StatusUnauthorized)
		return
	}

	var patchRequest UserPatchRequest
	if err := common.JsonParse(w, r, &patchRequest); err != nil {

		log.Printf("ERROR: updateUserHandler - Failed to parse JSON: %v", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	updatedProfile, err := updateUserAvatarAndDescriptionPostgres(ctx, requesterUserID, patchRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("INFO: User with ID %d not found for update.", requesterUserID)
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err.Error() == "no fields provided for update (user_img or description)" {
			http.Error(w, "No valid fields (user_img or description) provided for update.", http.StatusBadRequest)
			return
		}
		log.Printf("ERROR: updateUserHandler - Failed to update user profile %d: %v", requesterUserID, err)
		http.Error(w, "Failed to update user profile due to server error.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	common.JsonEncode(w, updatedProfile)
	log.Printf("INFO: User %d profile updated successfully.", idVal)
}
