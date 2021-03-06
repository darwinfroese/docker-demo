import React, { Component } from "react";
import axios from 'axios';
import CurrencyFormat from 'react-currency-format';
import './Login.css';

class Shop extends Component {
	constructor(props) {
		super(props);

		this.state = {
			items: []
		};
	}

	componentDidMount() {
		axios.get("/api/v1/shop").then(response => {
			console.log("response: ", response.data);
			this.setState({
				items: response.data
			})
		})
	}

	render() {
		const itemElements = [];

		this.state.items.forEach( item => {
			itemElements.push(
				<div className="item-container">
					<span className="item-title">{item.Name}</span>
					<span className="item-description">{item.Description}</span>
					<CurrencyFormat value={item.Price} displayType={'text'} thousandSeparator={true}
						renderText={value => <span className="item-price">{value}</span>} />
				</div>);
		});

		return (
			<div className="shop-container">
				{ itemElements }
			</div>
		);
	}
}

export default Shop;
