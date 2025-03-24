## 21st March 2025

# GIT AND WORKFLOW
This file is basically for learning Git.

# Readme features
Here are some features you can introduce to a README file using Markdown:
##1. Headers

    Use #, ##, or ### to create headings and subheadings for better organization.

    text
    # Project Title
    ## Subheading
    ### Smaller Subheading

##2. Lists

    Unordered Lists: Use -, *, or + for bullet points.

text
- Item 1
- Item 2

Ordered Lists: Use numbers followed by a period.

    text
    1. Step One
    2. Step Two

##3. Links

    
## To create hyperlinks on readme
To create hyperlinks ion readme use the squre brackets defining the name of the link follwed by the hyperlink in parenthesis.
### Example [Link Text](https://example.com)
There are various hyperlinks and each has a different format for creation:
###a. Basic Markdown Links

    Standard syntax for links:

text
[Link Text](https://example.com)

Example:

    text
    [Visit Google](https://www.google.com)

### b. Linking to Sections Within the README

    Use anchors (#) to link to specific sections of the README:

text
[Go to Features Section](#features)

Ensure the target heading exists in the README:

    text
    ## Features

### c. Linking to Files in the Repository

    Link to other files or folders in your project:

    text
    [Documentation](./docs/README.md)

    Use relative paths for files within the same repository.

### d. Email Links

    Use mailto: for email addresses:

    text
    [Contact Us](mailto:example@example.com)

### e. Image Hyperlinks

    Combine an image with a hyperlink:

text
[![Alt Text](image-url)](https://example.com)

Example:

    text
    [![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://ci.example.com)

### f. Plain Text URLs

    To display a URL as plain text without making it clickable, wrap it in backticks:

    text
    `https://example.com`

### g. HTML Links

    Use HTML tags for more customization (e.g., opening links in a new tab):

    xml
    <a href="https://example.com" target="_blank">Open Link</a>

### h. Advanced Markdown with Descriptions

    Provide context around the link:

    text
    For more details, see the [official documentation](https://docs.example.com).

### i. Badges with Links

    Add badges that link to external resources (e.g., CI/CD status, license):

    text
    [![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://ci.example.com)

### j. Linking to External Resources

    Include links to APIs, tools, or other related projects:

text
[API Documentation](https://api.example.com/docs)


   

## 4. Images

    Include screenshots or diagrams.

    text
    ![Alt Text](image-url)

## 5. Code Blocks

    Highlight code snippets using backticks.

text
undefined

print("Hello, World!")

    text
    undefined

## 6. Tables

    Create tables for structured data.

    text
    | Column A | Column B |
    |----------|----------|
    | Data A1  | Data B1  |
    | Data A2  | Data B2  |

## 7. Task Lists

    Add checklists for project milestones.

    text
    - [x] Completed Task
    - [ ] Pending Task

## 8. Emphasis

    Use * or _ for italic text and ** for bold text.

    Combine both for bold and italic text.

## 9. Horizontal Lines

    Separate sections with lines using ---, ***, or ___.

    text
    ---

## 10. Strikethrough

    Strike out text using ~~.

    text
    ~~This text is crossed out~~

## 11. Anchors

    Create links to specific sections within the README.

    text
    [Go to Subheading](#subheading)

## 12. Math Expressions

    Render mathematical formulas using LaTeX-style syntax (if supported).

By combining these features, you can create a visually appealing and highly informative README file!



## 22nd March 2025
# GIT BRANCHING
*A branch is basically a separate line of development of the codebase.

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

### Merging remotely
Go to your Git website host e.g GitHub,GitLab,BitBucket... and make a pull request(PR),follow the prompts to merge 

the master and the feature branch contents.

### Merging locally
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


## 23rd March 2025

## Resetting or reverting staging and commit in Git(Undoing Git)
### Resetting
### a.Staging
When you add changes(staging) ready to be commited then you realize they are not the right changes you intended to stage then you can simply reverse this process using ### git reset

### Workflow
-git add filename

-git reset / git reset filename

### b.commit
When you commit changes already staged ready to be pushed then you realize they are not the right changes you intended to commit then ypu can simply reverse this process to the initial status before staging using ### git reset HEAD~1

### Workflow
-git add filename

-git commit -m "message"

-git reset HEAD~1(This will unstage and uncommit the changes made)

### Alternatively
We can use the git log to obtain the commit hashes which can also be used to reverse the changes made on a file.

N/B:-To scroll the git log list use the space bar or the enter button.

Copy the commit hash and use git reset #hashcode (This uncommit and unstage the changes made to the file but will not do away with them)

To completely do away with the commits and staging use git reset --hard #hashcode


### Reverting
 Purpose of git revert: It is designed to undo committed changes by creating a new commit that reverses the effects of a previous commit. It does not affect uncommitted or staged changes.

Scope: git revert works exclusively on commits, not on the staging area or untracked files.

This is helpful to restore already pushed changes.
