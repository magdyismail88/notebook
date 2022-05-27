



exports.flashError = (msg, vm) => {
  setTimeout(() => {
    vm.status = true;
    vm.isError = true;
    vm.msg = msg;
    setTimeout(() => {
      vm.status = false;
      vm.isError = false;
      vm.msg = '';
    }, 2000);
  });
}

exports.flashSuccess = (msg, vm) => {
  setTimeout(() => {
    vm.status = true;
    vm.isError = false;
    vm.msg = msg;
    setTimeout(() => {
      vm.status = false;
      vm.msg = '';
    }, 2000);
  });
}


exports.center = (id) => {
  // obtain the object reference for the textarea>
  var txtarea = document.getElementById(id);
  // obtain the index of the first selected character
  var start = txtarea.selectionStart;
  // obtain the index of the last selected character
  var finish = txtarea.selectionEnd;
  //obtain all Text
  var allText = txtarea.value;
  
  // obtain the selected text
  var sel = allText.substring(start, finish);
  //append te text;
  var newText=allText.substring(0, start)+"<center>"+sel+"</center>"+allText.substring(finish, allText.length);
  
  txtarea.value=newText;
  console.log(newText);
  // do something with the selected content
}


// export default class html {

//   static Center(id) {
//     // obtain the object reference for the textarea>
//     var txtarea = document.getElementById(id);
//     // obtain the index of the first selected character
//     var start = txtarea.selectionStart;
//     // obtain the index of the last selected character
//     var finish = txtarea.selectionEnd;
//     //obtain all Text
//     var allText = txtarea.value;
    
//     // obtain the selected text
//     var sel = allText.substring(start, finish);
//     //append te text;
//     var newText=allText.substring(0, start)+"<center>"+sel+"</center>"+allText.substring(finish, allText.length);
    
//     txtarea.value=newText;
//     console.log(newText);
//     // do something with the selected content
//   }


// }


// document.getElementById("btnCenter").addEventListener("click",getSel);