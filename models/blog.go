package models

type Blog struct {
    ID               string `bson:"_id,omitempty" json:"id"`
    IntroductionText string `bson:"introduction_text" json:"introduction_text"`
    IntroductionImage string `bson:"introduction_image" json:"introduction_image"`

    ChallengeText string `bson:"challenge_text" json:"challenge_text"`
    ChallengeImage string `bson:"challenge_image" json:"challenge_image"`

    SolutionText string `bson:"solution_text" json:"solution_text"`

    ConclusionText string `bson:"conclusion_text" json:"conclusion_text"`
    ConclusionImage string `bson:"conclusion_image" json:"conclusion_image"`

    PreferredURL string `bson:"preferred_url" json:"preferred_url"`
}