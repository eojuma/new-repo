# New file
This file is basically for learning Git.

I will make short notes here on how to use git to keep track of every changes that you make on your files.

This will be very appealing to my fellow Git beginners.

## 21st March 2025
### Branching in Git
*A branch is basically a separate line of development of the codebase.

*I have learnt how to create git branches,how to add,commit and merge to the main branch.

### Creating a new branch
To create a new branch;use git checkout -b branchname--->This creates a new branch and immediately 

switches to the newly created branch
### Listing all available branches on the current directory
To check if the created branch is available;git branch,this will list all the branches on your current directory.

To know which branching you're working on,after listing all the available branches,check the one with an asterisk(*)
### Switching to an existing branch
To switch to an existing branch;git checkout branchname
### Git Branching workflow
While on the main/master branch,switch to the feature branch created using(git checkout branchname).

Make whatever developments you're up to.

-Git add filename

-Git commit -m "updates"

Git push -u origin feature-branchname

### Git workflow
Git add filename

Git commit -m "message"

Git push

NOTE:-Once a file has been commited,any changes that will be made later can simply be added and commited simulteneouly;

We use git commit -am "message"---->this action both adds and commits the changes made(does not work for new files)

Then push the changes to the remote repository. 

### Merge remotely
Go to your Git website host e.g GitHub,GitLab,BitBucket... and make a pull request(PR),follow the prompts to merge 

the master and the feature branch contents.

### Merge locally
To merge locally,switch to the feature-branch from where yo want to make the changes.

Make the necessary changes,add and commit while on the same branch.

When pushing from the feature-branch(git push -u origin feature-branch) you'll be guided on how to merge remotely.

To merge locally,git merge maser/main branch,switch to the main branch and git pull to see the changes made.

### Deleting the feature branch
Once the two branches are merged,locally delete the feature branch using;git branch -d feature-branchname

This can also be remotely,immediately after a successful merge,you will be asked if you want delete the feature branch.

### Update the main branch
To see the changes made after merging the two branches,git pull, while on the main branch.

Before deleting the feature branch we can also see the changes made on the main branch using ~git diff feature-branchname~

At this point remember the changes have not been made remotely and the feature branch is still existing.

### All commits
To have access to all commits to a given repo since existence,use git log

[GitHub](https://github.com)

[Emailaddress](mail:evansodhiambo658@gmail.com)

[Email Me](mailto:evansodhiambo658@gmail.com)




