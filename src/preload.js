const {
  contextBridge,
  ipcRenderer
} = require("electron");

const { HardwareMonitorClient } = require("./proto/service_grpc_web_pb");
const {EmptyRequest } = require("./proto/service_pb");


window.addEventListener('DOMContentLoaded', () => {
  const replaceText = (selector, text) => {
    const element = document.getElementById(selector)
    if (element) element.innerText = text
  }

  for (const dependency of ['chrome', 'node', 'electron']) {
    replaceText(`${dependency}-version`, process.versions[dependency])
  }


  // seems to break after a few seconds: https://github.com/ngx-grpc/ngx-grpc/issues/41
  
  setTimeout(()=>{
    var client = new HardwareMonitorClient('http://localhost:8080');


    var request = new EmptyRequest();
    // Dont worry about the empty Metadata for now, thats covered in another article :)
    var stream = client.monitor(request, {});
    // Start listening on the data event, this is the event that is used to notify that new data arrives
    stream.on('data', function (response) {
      // Convert Response to Object
      var stats = response.toObject();
      // // Set our variable values
      // setCPU(stats.cpu);
      // setMemoryFree(stats.memoryFree);
      // setMemoryUsed(stats.memoryUsed);
      replaceText(`electron-native-addon`, stats.cpu)
    });
  }, 1000);
  



  ipcRenderer.invoke("a", "b").then((x) => {
    replaceText(`electron-native-addon`, x)
  })

})

