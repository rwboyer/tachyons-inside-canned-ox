import React from "react";
import ReactDOM from "react-dom";
import Header, {List, ListItem, TachLikeButton} from "./components/header"
import css from "./index.css";
//import {FormWizard} from "./components/form"
import FormWizard from "./components/form"
import App from "./components/form-alt"
import ReactCompareImage from 'react-compare-image';
import Gallery from 'react-photo-gallery';

const Index = () => {
  return <div class="pv5">Hello React!</div>;
};

const items = [
  {key: 'item1', isDone: false, text: 'Listing One'},
  {key: 'item2', isDone: false, text: 'Listing Two'},
  {key: 'item3', isDone: false, text: 'Listing Three'},
  {key: 'item4', isDone: false, text: 'Listing Four'},
  {key: 'item5', isDone: false, text: 'Listing Five'},
]

ReactDOM.render(
  <div>
  <Header />
  <TachLikeButton />
  <List title='Todo'>
  {items.map(item => (
    <ListItem
      key={item.key}
      isDone={item.isDone}
      children={item.text} />
  ))}
  </List>
  </div>
, document.getElementById("react-target"));

ReactDOM.render(
  <App />
  , document.getElementById("formik-target")
);

var img_holder = document.getElementById('#img-list');
var img_list = img_holder.getElementsByTagName('img');
var arr = [].slice.call(img_list);
var img_simple = arr.map(a => a.src);
console.log(img_simple);
console.log(img_list);


ReactDOM.render(
  <div style={{ maxWidth: '100%' }}>
  <ReactCompareImage
    leftImage={img_list[0].src}
    rightImage={img_list[1].src}
    sliderLineWidth={2}
    handleSize={50}
  />
  </div>
  , document.getElementById("image-target")
);

var port_holder = document.getElementById('port');
var port_list = port_holder.getElementsByTagName('dd');
var ddarr = [].slice.call(port_list);
var port_simple = ddarr.map(a => a.getAttribute('data-item'));
console.log(port_simple);
console.log(port_list);

var port_display = port_simple.map(i => {return {src: i, width:4, height: 6}});

ReactDOM.render(
  <Gallery photos={port_display}
    columns={4}
  />
  , document.getElementById("port-target")
);
