# Minimal Implementation Of Langauge Server Protocol in Go

This is a barebones implementation of a Language Server Protocol (LSP) in Go, built to get familiar with LSP development. It supports fundamental features like hover information, go-to definitions, code actions, and autocomplete. The project is aimed at learning LSP concepts and basic Go development."

## How To Install:
To install, just clone the repo and run go build main.go.

![Command to clone the repo](https://raw.githubusercontent.com/WizardOfOz-1/Golang-LSP/refs/heads/main/assets/git_clone.png)

## How To Use:

Personally I Use Neovim, and I do not know the implementation process in other editors. So I will just lay out the steps for Neovim. 

![Creating The LSP Definition](https://raw.githubusercontent.com/WizardOfOz-1/Golang-LSP/refs/heads/main/assets/start_client.png)
(**note the cmd field should point to combiled binary **)

->  And an autocommand to trigger the LSP whenever a markdown file is opened:

![Triggering The LSP](https://raw.githubusercontent.com/WizardOfOz-1/Golang-LSP/refs/heads/main/assets/autocmd.png)

Here I've added some keybindings just so its easier to do things such as <kbd>codeactions</kbd>,<kbd>Hovers</kbd> and <kbd>Goto Definitions</kbd>. For Autocomplete, <kbd>CTRL + X</kbd><kbd>CTRL+O</kbd> triggers Neovim's built in auto complete so we can use that. This will happen automatically if you're using a completion engine like say [nvim-cmp](https://github.com/hrsh7th/nvim-cmp). 

# Screenshots:
# Hover:
<kbd> SPACE gd</kbd> will trigger hover definition. (assuming leader is mapped to spacebar)
![textDocument/hover](https://raw.githubusercontent.com/WizardOfOz-1/Golang-LSP/refs/heads/main/assets/screenshot-20241021111411.png)

# Diagnostics:
Upon an instance of the word "VS Code" in the current buffer, the LSP will send out an **ERROR** diagnostic. (lol)
![textDocument/hover](https://github.com/WizardOfOz-1/Golang-LSP/blob/main/assets/screenshot-20241021111503.png?raw=true)

# Code Actions:
Upon a diagnostic error (in our case the word "VS Code") we can perform code actions, with <kbd>SPACE ca</kbd> (assuming leader is mapped to spacebar)
![textDocument/hover](https://github.com/WizardOfOz-1/Golang-LSP/blob/main/assets/screenshot-20241021111514.png?raw=true)
The user can pick any one of the above to either censor the word "VS Code" or to replace it with a superior editor (**NEOVIM BTW**)

### Case Picking 1:
___
![Code Action To Replace Neovim With Superior Editor](https://github.com/WizardOfOz-1/Golang-LSP/blob/main/assets/screenshot-20241021111612.png?raw=true)<br>
If user chooses 1, "VS Code" is replaced with Neovim, and the LSP affirms your decision. (lol)


### Case Picking 2:
___
![Code Action To censor VS Code](https://github.com/WizardOfOz-1/Golang-LSP/blob/main/assets/screenshot-20241021111528.png?raw=true)<br>
If user chooses 2, "VS Code" is censored to "VS C*de) (lol)
___
In case there are no diagnostics, codeactions wont trigger, and the editor will let the user know.<br>
![No Code Actions](https://github.com/WizardOfOz-1/Golang-LSP/blob/main/assets/screenshot-20241021111434.png?raw=true)

# Autocomplete:
Using <kbd>CTRL X</kbd><kbd>CTRL O</kbd> will trigger autocomplete in Neovim, upon which the editor will receive autocomplete suggestion from the LSP.
![AutoComplete](https://github.com/WizardOfOz-1/Golang-LSP/blob/main/assets/screenshot-20241021111658.png?raw=true)
( This happens automatically if you're using a completion engine as nvim-cmp )

