# New file
This file is basically for learning Git.
I will make short notes here on how to use git to keep track of every changes that you make on your files.
This will be very appealing to my fellow Git beginners.

# 21st March 2025
# Branching in Git
A branch is basically a separate line of development of the codebase.
I have learnt how to create git branches,how to add,commit and merge to the main branch.

# Creating a new branch
To create a new branch;use git checkout -b branchname--->This creates a new branch and immediately 
switches to the newly created branch
# Listing all available branches on the current directory
To check if the created branch is available;git branch,this will list all the branches on your current directory.
To know which branching you're working on,after listing all the available branches,check the one with an asterisk(*)
# Switching to an existing branch
To switch to an existing branch;git checkout branchname
# Git Branching workflow
While on the main/master branch,switch to the feature branch created using(git checkout branchname).
Make whatever developments you're up to.
Git add filename
Git commit -m "updates"
Git push -u origin feature-branchname

# Merge remotely
Go to your Git website host e.g GitHub,GitLab,BitBucket... and make a pull request(PR),follow the prompts to merge 
the master and the feature branch contents.

# Deleting the feature branch
Once the two branches are merged,locally delete the feature branch using;git branch -d feature-branchname
This can also be remotely,immediately after a successful merge,you will be asked if you want delete the feature branch.

# Update the main branch
To see the changes made after merging the two branches,git pull, while on the main branch.
Before deleting the feature branch we can also see the changes made on the main branch using ~git diff feature-branchname~
At this point remember the changes have not been made remotely and the feature branch is still existing.

# All commits
To have access to all commits to a given repo since existence,use git log