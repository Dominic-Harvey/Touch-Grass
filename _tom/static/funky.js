let x;
let y;
let x_speed;
let y_speed;

let dvd;

function preload() {
  newImage();
}

function setup() {
  createCanvas(windowWidth*0.9, windowHeight*0.9);
  x = random(width);
  y = random(height);
  x_speed = 5;
  y_speed = 5;
}

function draw() {
  background(0);
  image(dvd, x, y, 80, 60);

  x += x_speed;
  y += y_speed;

  if (x + 80 >= width || x <= 0) {
    x_speed *= -1;
    newImage();
  }
  if (y + 60 >= height || y <= 0) {
    y_speed *= -1;
    newImage();
  }
}

function newImage() {
  fetch('/images')
    .then(response => response.blob())
    .then(blob => {
      dvd = createImg(URL.createObjectURL(blob));
      dvd.hide();
    })
    .catch(error => {
      console.error('Error fetching new image:', error);
    });
}