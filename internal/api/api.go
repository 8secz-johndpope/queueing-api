package api

import (
	"encoding/json"
	"gitlab.com/projectreferral/queueing-api/configs"
	"gitlab.com/projectreferral/queueing-api/internal/mq"
	"gitlab.com/projectreferral/queueing-api/models"
	"log"
	"net/http"
	"os"
)

func Init(){
	configs.BrokerUrl = os.Getenv("BROKERURL")
	log.Println(configs.BrokerUrl)
	mq.CreateFailedMessageQueue()
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	if mq.TestQ(w) {
		w.WriteHeader(http.StatusOK)
	}
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
	queue := models.QueueDeclare{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&queue)
	if !HandleError(err, w) {
		mq.RabbitCreateQueue(w,queue,false)
	}
}

func CreateExchange(w http.ResponseWriter, r *http.Request) {
	exchange := models.ExchangeDeclare{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&exchange)
	if !HandleError(err, w) {
		mq.RabbitCreateExchange(w,exchange)
	}
}

func BindExchange(w http.ResponseWriter, r *http.Request) {
	bind := models.QueueBind{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&bind)
	if !HandleError(err, w) {
		mq.RabbitQueueBind(w,bind)
	}
}

func PublishToExchange(w http.ResponseWriter, r *http.Request) {
	publish := models.ExchangePublish{}
	err := json.NewDecoder(r.Body).Decode(&publish)
	if !HandleError(err, w) {
		mq.RabbitPublish(w,publish)
	}
}

func ConsumeQueue(w http.ResponseWriter, r *http.Request) {
	consume := models.QueueConsume{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&consume)
	if !HandleError(err, w) {
		mq.RabbitConsume(w,consume)
	}
}

func SubscribeQueue(w http.ResponseWriter, r *http.Request) {
	subscribe := models.QueueSubscribe{
		MaxRetry: -1,   //default limit is none
		Arguments: nil,
	}
	err := json.NewDecoder(r.Body).Decode(&subscribe)
	if !HandleError(err, w) {
		mq.RabbitSubscribe(w,subscribe)
	}
}

func UnSubscribeQueue(w http.ResponseWriter, r *http.Request) {
	subId := models.QueueSubscribeId{}
	err := json.NewDecoder(r.Body).Decode(&subId)
	if !HandleError(err, w) {
		mq.RabbitUnsubscribe(subId.ID)
	}
}

func MessageAck(w http.ResponseWriter, r *http.Request){
	acknowledge := models.MessageAcknowledge{}
	err := json.NewDecoder(r.Body).Decode(&acknowledge)
	if !HandleError(err, w) {
		if acknowledge.GetID() != "" {
			mq.RabbitAck(w,acknowledge)
			return
		}
		w.WriteHeader(403)
	}
}

func MessageReject(w http.ResponseWriter, r *http.Request){
	reject := models.MessageReject{}
	err := json.NewDecoder(r.Body).Decode(&reject)
	if !HandleError(err, w) {
		if reject.GetID() != "" {
			mq.RabbitReject(w,reject)
			return
		}
		w.WriteHeader(403)
	}
}

func DumpData(w http.ResponseWriter, r *http.Request) {
	dataUser := dataUser{}
	err := json.NewDecoder(r.Body).Decode(&dataUser)
	if !HandleError(err, w) {
		mq.ArrayDump(w,dataUser.Password)
	}
}

type dataUser struct {
	Password       string     `json:"password"`
}