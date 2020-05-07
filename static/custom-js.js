
function closeApp() {
   window.fetch(window.location.host+"/quit")
   window.close();
}
