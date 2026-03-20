# MicTray

MicTray is a lightweight, natively compiled macOS menu bar application designed to give you peace of mind and total control over your microphone. 

It provides instant visual feedback—showing a red mic when your microphone is active, and a gray mic when your volume is safely set to zero. You can effortlessly mute and unmute your mic across all applications globally with just a single click.

## Prerequisites

To build and run MicTray, you will need:
- **macOS**
- **Go 1.20+** installed correctly on your system (e.g. `brew install go`)

## Building and Installing

The project includes a Makefile to automatically compile the Go binary and package it into a native macOS Application Bundle (`MicTray.app`) that runs silently in the background without cluttering your OS Dock.

1. **Clone the repository**  
   ```bash
   git clone https://github.com/tiagodll/mictray.git
   cd mictray
   ```

2. **Build the Apple Application Bundle**  
   To build the Go executable and bundle it into `MicTray.app`, simply run:
   ```bash
   make
   ```

3. **Install to your Applications folder**  
   After running `make`, a `MicTray.app` folder will be generated in your project root. To install it globally, simply move it into your Mac's `/Applications/` folder:
   ```bash
   mv MicTray.app /Applications/
   ```
   *You can then launch it from Launchpad or Finder like any normal macOS application.*

## Usage

- **Left-Click the Tray Icon**: Instantly toggles the system input volume between 0 (Mute) and 100 (Hot).
- **Context Menu**: Click the menu items or right-click to access the "About" page or cleanly "Quit" the application.
