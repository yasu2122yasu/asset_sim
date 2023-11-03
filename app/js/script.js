function inputChange(event) {
  msg.innerText = '音量は ' + volumeSlider.value + ' です';
}

let volumeSlider = document.getElementById('volumeSlider');
volumeSlider.addEventListener('input', inputChange);
let msg = document.getElementById('msg');
