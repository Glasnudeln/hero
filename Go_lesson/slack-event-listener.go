package main

import (
  "fmt"
)

func (s *SlackListener) ListenAndResponse() {
  // Start listening slack event
  rtm := s.client.NewRTM()
  go rtm.ManageConnection()

  // Handle slack events
  for msg := range rtm.IncomingEvents {
    swtich ev := msg.Data.(type) {
    case *slack.MessageEvent:
      if err := s.handleMessageEvent(ev); err != nul {
        log.Printf("[ERROR] Failed to handle message: %s", err)
      }
    }
  }
}

func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error

var attachment = slack.Attachment {
  Text:       "Which book do you want? :book:",
  Color:      "#f9a41b",
  CallbackID: "book",
  Actions: []slack.AttachmentAction{
    {
      Name: actionSelect,
      Type: "select",
      Options: []slack.AttachmentActionOption{
        {
          Text:   "GoIntroduction",
          Value:  "GoIntroduction",
        },
        {
          Text:   "RubyIntroduction",
          Value:  "RubyIntroduction",
        },
        {
          Text:   "PerlIntroduction",
          Value:  "PerlIntroduction",
        },
        {
          Text:   "vue.jsIntroduction",
          Value:  "vue.jsIntroduction",
        },
        {
          Text:   "JavaIntroduction",
          Value:  "JavaIntroduction",
        },
      },
    },

    {
      Name:   actionCancel,
      Text:   "Cancel",
      Type:   "button",
      Style:  "danger",
    },
  },
}
