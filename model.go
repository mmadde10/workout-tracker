package main

// AppInfo sturct that defines info of app
type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type UserModel struct {
	Id           uint   `json:"id"`
	Username     string `json:username"`
	Email        string `json:email"`
	Bio          string `json:bio`
	PasswordHash string `gorm:"column:password;not null"`
}

type Workout struct {
	Name      string     `json:"name"`
	StartTime string     `json:"startTime"`
	EndTime   string     `json:"endTime"`
	Exercises []Exercise `json:"exercise"`
	StartedBy UserModel  `json:"startedBy"`
}

type Exercise struct {
	Name                string `json:"name"`
	ExerciseType        string `json:"exerciseType"`
	EquipmentRequired   string `json:"equipmentRequired"`
	TargetedMuscleGroup string `json:"targetedMuscleGroup"`
	RecommendedSets     int    `json:"recommendedSets"`
	RecommendedReps     int    `json:"recommendedSets"`
	Sets                []Set  `json:"sets"`
}

type Set struct {
	Weight    int `json:"weight"`
	SetNumber int `json"setNumber"`
	Reps      int `json:"reps"`
}
