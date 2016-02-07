# xmppsend

This is a (very) small CLI tool to send a single message via XMPP.

I use it to deliver notifications from [Zabbix](http://www.zabbix.com/) because the Jabber client in there is broken TLS-wise.

## Usage

Install go and then run `go install github.com/xperimental/xmppsend`, then you should have it available in your path.

Then use:

```bash
$ xmppsend -help
Usage of xmppsend:
  -msg string
    	Message to send.
  -password string
    	Password of sender.
  -to string
    	JID of recipient.
  -user string
    	JID of sender.
```
