package core

import (
	"errors"
	"time"

	sqs "github.com/tommzn/aws-sqs"
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	utils "github.com/tommzn/go-utils"
	model "github.com/tommzn/recipeboard-core/model"
)

// Creates a new recipe message for goven recipe.
// Will assign a new id and sequence number.
func newRecipeMessage(recipe model.Recipe, action model.RecipeAction) model.RecipeMessage {

	return model.RecipeMessage{
		Id:       utils.NewId(),
		Action:   action,
		Sequence: time.Now().UnixNano(),
		Recipe:   recipe,
	}
}

// newSqsPublisher creates publisher to send messages to AWS SQS.
func newSqsPublisher(conf config.Config, logger log.Logger) model.MessagePublisher {
	queueName := conf.Get("aws.sqs.queue", nil)
	return &sqsPublisher{
		client:    sqs.NewPublisher(conf),
		queueName: queueName,
		logger:    logger,
	}
}

// Send wll publish given message on an AWS SQS queue.
func (publisher *sqsPublisher) Send(message model.RecipeMessage) error {

	if publisher.queueName == nil {
		return errors.New("Missing AWS SQS queue to publish messages for recipe actions.")
	}

	sqsId, err := publisher.client.Send(message, *publisher.queueName)
	if err == nil {
		publisher.logger.Infof("Message for recipe %s, action %s published to sqs as %s", message.Id, message.Action, *sqsId)
	}
	return err
}
