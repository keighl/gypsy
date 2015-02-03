package api

import (
  m "github.com/keighl/gypsy-cab/models"
  u "github.com/keighl/gypsy-cab/utils"
  "github.com/keighl/mandrill"
  r "github.com/dancannon/gorethink"
)

var  (
  Config *u.Configuration
  DB *r.Session
)

//////////////////////////////
// API DATA //////////////////

type Data struct {
  APIToken string `json:"api_token,omitempty"`
  *Error `json:"error,omitempty"`
  *Message `json:"message,omitempty"`
  *m.User `json:"user,omitempty"`
  *m.PasswordReset `json:"password_reset,omitempty"`
  *m.Item `json:"item,omitempty"`
}

// For older clients that may expect a data.{} envelope
type LegacyEnvelope struct {
  *Data `json:"data"`
}


//////////////////////////////
// API MESSAGE ///////////////

type Message struct {
  Message string `json:"message,omitempty"`
}

//////////////////////////////
// API ERROR /////////////////

type Error struct {
  Message string `json:"message,omitempty"`
  Details []string `json:"details,omitempty"`
}

func ServerErrorEnvelope() (Data) {
  data := Data{}
  data.Error = &Error{"There was an unexpected error!", []string{}}
  return data
}

func ErrorEnvelope(message string, details []string) (Data) {
  data := Data{}
  data.Error = &Error{message, details}
  return data
}

func MessageEnvelope(message string) (Data) {
  data := Data{}
  data.Message = &Message{message}
  return data
}

//////////////////////////////
// MAILER ////////////////////

var sendEmail = func(message *mandrill.Message) (bool) {
  client := mandrill.ClientWithKey(Config.MandrillAPIKey)
  _, Error, err := client.MessagesSend(message)
  if (Error != nil || err != nil) { return false }
  return true
}
