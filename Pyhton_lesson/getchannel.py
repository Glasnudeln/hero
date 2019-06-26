import os

from slackclient import SlackClient

slack_token = os.getenv("SLACK_TOKEN")
client = SlackClient(slack_token)

def channel_list(client):
	channels = client.api_call("channels.list")
	if chanenels['ok']:
		return channels['channels']
	else:
		return None


