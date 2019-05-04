package models

type MessageAnswer struct {
    Status  int    `json:"status, int" example:"100"`
    Message string `json:"message, string" example:"ok"`
}

type AddAreaAnswer struct {
    Status  int               `json:"status, int" example:"101"`
    Message string            `json:"message, string" example:"user found"`
    Data    AddAreaAnswerData `json:"data"`
}

type TrainAnswer struct {
    Status  int             `json:"status, int" example:"101"`
    Message string          `json:"message, string" example:"user found"`
    Data    TrainAnswerData `json:"data"`
}

func GetSuccessAnswer(message string) *MessageAnswer {
    return &MessageAnswer{
        Status:  100,
        Message: message,
    }
}

func GetAddAreaAnswer(data *AddAreaAnswerData) *AddAreaAnswer {
    return &AddAreaAnswer{
        Status:  101,
        Message: "ok",
        Data:    *data,
    }
}

func GetTrainAnswer(data *TrainAnswerData) *TrainAnswer {
    return &TrainAnswer{
        Status:  102,
        Message: "ok",
        Data:    *data,
    }
}

func GetErrorAnswer(error string) *MessageAnswer {
    return &MessageAnswer{
        Status:  200,
        Message: error,
    }
}

var IncorrectJsonAnswer = MessageAnswer{
    Status:  201,
    Message: "incorrect JSON",
}
