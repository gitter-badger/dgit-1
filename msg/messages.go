package msg

var Welcome = `
Welcome to dgit!

A private key has been generated for you and is stored in {{.keyringProvider}}.

Below is your dgit public address, share this with others so they can grant you write access to dgit repos.
{{.userAddress}}
`

var AddDgitToRemote = `
dgit would like to add {{.repourl}} to the '{{.remote}}' remote. This allows 'git push' to mirror this repository to dgit.
`

var AddDgitToRemoteConfirm = `Is that ok?`

var AddedDgitToRemote = `
Success, dgit is now superpowering the '{{.remote}}' remote.
Continue using your normal git workflow and enjoy being decentralized.
`

var AddDgitRemote = `
dgit would like to add the '{{.remote}}' remote to this repo so that you can fetch directly from dgit.
`

var AddDgitRemoteConfirm = AddDgitToRemoteConfirm

var AddedDgitRemote = `
Success, dgit is now accessible under the '{{.remote}}' remote.
'git fetch {{.remote}}' will work flawlessly from your decentralized repo.
`

var FinalInstructions = `
You are setup and ready to roll with dgit.
Just use git as you usually would and enjoy a fully decentralized repo.

If you would like to clone this dgit repo on another machine, simply run 'git clone {{.repourl}}'.

If you use GitHub for this repo, we recommend adding a dgit action to keep your post-PR branches in sync on dgit.
You can find the necessary action here:
https://github.com/quorumcontrol/dgit-github-action

Finally for more docs or if you have any issues, please visit our github page:
https://github.com/quorumcontrol/dgit
`

var PromptRepoNameConfirm = `It appears your repo is '{{.repo}}', is that correct?`

var PromptRepoName = `What is your full repo name?`

var PromptRepoNameInvalid = `Enter a valid repo name in the form '${user_or_org}/${repo_name}'`

var PrivateKeyNotFound = `
Could not load your dgit private key from {{.keyringProvider}}. Try running 'dgit init' again.
`

var RepoCreated = `
Your dgit repo has been created at '{{.repo}}'.

dgit repo identities and authorizations are secured by Tupelo - this repo's unique id is:
'{{.did}}'.

Storage of the repo is backed by Sia Skynet.
`

var RepoNotFound = `
dgit repository does not exist.

You can create a dgit repository by running 'dgit init'.
`

var RepoNotFoundInPath = `
No .git directory found in {{.path}}.

Please change directories to a git repo and run '{{.cmd}}' again.

If you would like to create a new repo, use 'git init' normally and run '{{.cmd}}' again.
`
