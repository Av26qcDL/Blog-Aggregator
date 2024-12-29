README.md for Blog Aggregator 

Introduction

Welcome to the Blog Aggregator, a CLI tool designed to streamline your experience in gathering and managing content from the web. It empowers users to:

Add RSS feeds from various sources for seamless collection
Store aggregated posts in a PostgreSQL database for easy access
Follow and unfollow RSS feeds curated by other users
Explore summarized posts directly in your terminal with links to the full articles
This tool is the ideal companion for anyone looking to keep up with their favorite blogs, news sites, podcasts, and more effortlessly. Whether you seek to track personal interests or gather professional insights, Blog Aggregator simplifies the process, all from the comfort of your terminal! 

Prerequisites

Before diving into the Blog Aggregator, ensure that you have the following installed on your system:

PostgreSQL: A robust database management system for storing your collected posts. (Recommended version: 14.15)
Go: The language of choice for this CLI tool, offering high performance and speed. (Recommended version: 1.23.0)
Please verify that your system meets these requirements. This tool is primarily developed for Windows, but it can also run on macOS and Linux with appropriate configurations. 

Installation Instructions

Follow these steps to get your environment ready for Blog Aggregator:

Install PostgreSQL

Visit the PostgreSQL download page.
Choose the installer that matches your operating system.
Run the installer and complete the setup wizard.
Ensure PostgreSQL is running by checking its status via your terminal or system service manager.
Note: For macOS, consider using package managers like Homebrew: brew install postgresql.

Install Go

Visit the Go installation guide.
Download and run the installer for your platform (Windows, macOS, or Linux).
Follow the steps provided in the wizard to complete the installation.
Check your installation by running go version in the terminal to confirm Go is set up correctly.
Once installed, you are ready to configure and run the Blog Aggregator tool! 

Configuration and Usage

Setting Up a Config File

To configure Blog Aggregator:

Navigate to the directory in your system where you wish to store your configuration file.
Open your terminal in that directory.
Use a text editor like Nano to create the config file:
nano myname.conf
Enter your configuration settings, such as database credentials, API keys, or other necessary parameters (example: connection strings).
Save and close the file with Nano (CTRL + X, then Y to confirm, and Enter).
Ensure your configuration file is correctly set up to allow your application to connect to necessary resources.

Usage

Here are some essential commands to wield Blog Aggregator:

Build the Application: Compile the code to create the executable.

go build -o blog_aggregator

Browse Feeds: View available RSS feeds directly in your terminal.

./blog_aggregator browse

Add an RSS Feed: Add a new feed to your collection using its URL.

./blog_aggregator addfeed <url>

Follow a Feed: Start following a specified feed using its ID.

./blog_aggregator follow <feed_id>

This guide should help you seamlessly start and operate the Blog Aggregator. 
 
Pushing Modifications to GitHub

After you've made changes to your Blog Aggregator project, follow these steps to push your modifications to GitHub:

Initialize or Navigate to Your Git Repository

If you're starting fresh with a new repository:

git init

If you're already in a Git-tracked project, navigate to its directory.
Add Your Changes

Stage all your changes for commit:

git add .

Commit Your Changes

Commit the staged changes with a descriptive message:

git commit -m "Add feature or fix: brief description"

Push to GitHub

If it's your first time pushing to the repository, you'll need to set the remote URL:

git remote add origin https://github.com/github-username/repo-name.git

Push your changes to the remote repository:

git push origin main

Make sure to replace github-username and repo-name with your actual GitHub username and the repository's name. These steps will ensure your modifications are shared and visible on GitHub.