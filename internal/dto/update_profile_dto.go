package dto

type UpdateProfileRequestDTO struct {
	UserProfile   *UserProfileDTO    `json:"user_profile" validate:"required"`
	UserInterest  []*UserInterestDto `json:"user_interest" validate:"required"`
	UserLifestyle *UserLifestyleDTO  `json:"user_lifestyle" validate:"required"`
	UserRoutine   *UserRoutineDTO    `json:"user_routine" validate:"required"`
}
