# YVRDev (#vandevs) Slack Invite App  

![yvrdev logo](https://avatars1.githubusercontent.com/u/17320429?s=200&v=4)

## Setup

1.) Clone this repo
```
git clone git@github.com:YVRDev/go-slack-invite.git && cd go-slack-invite
```

2.) Add your [slack token](https://api.slack.com/custom-integrations/legacy-tokens) to a `.secrets` file
```
echo token=<my-slack-token-goes-here> > .secrets
```  
  
3.) Push to zeit.co using [now](https://zeit.co/now) and add the token as an env var
```
now -e $(cat .secrets)
```


  

