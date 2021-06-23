const path = require('path')
const blackfriday = require('./gonodemodule/golang');


const { app, BrowserWindow } = require('electron')

function createWindow() {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js')
    }
  })

  win.loadFile('index.html')

  console.log(blackfriday.hello())
}
  app.whenReady().then(() => {
    createWindow()

    app.on('activate', function () {
      if (BrowserWindow.getAllWindows().length === 0) createWindow()
    })
  })


  app.on('window-all-closed', function () {
    if (process.platform !== 'darwin') app.quit()
  })

