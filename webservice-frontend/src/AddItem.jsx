import axios from 'axios';
import React, { Component } from "react";

import './Login.css';

class AddItem extends Component {
	postItem() {
		let name = document.getElementById("item-name-input").value;
		let desc = document.getElementById("item-desc-input").value;
		let price = document.getElementById("item-price-input").value;

		axios.post("/api/v1/items", {
			Name: name, Description: desc, Price: price
		});

		document.getElementById("item-name-input").value = "";
		document.getElementById("item-desc-input").value = "";
		document.getElementById("item-price-input").value = "";
	}

	render () {
		return (
			<div className="add-container">
				<input type="text" placeholder="name" id="item-name-input"></input>
				<input type="text" placeholder="description" id="item-desc-input"></input>
				<input type="text" placeholder="price" id="item-price-input"></input>
				<div className="add-item-button" onClick={this.postItem}>Add Item</div>
			</div>
		);
	}
}

export default AddItem;
