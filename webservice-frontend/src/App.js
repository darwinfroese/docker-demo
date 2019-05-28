import React from 'react';
import { Link, Route, Switch } from "react-router-dom";
import Login from "./Login";
import Shop from "./Shop";
import AddItem from "./AddItem";
import './App.css';

function App() {
  return (
    <div className="app">
			<div className="header-container">
				<header className="app-header"> THE SHOP </header>
				<div className="nav-container">
					<Link to="/login" className="nav-item">Login</Link>
					<Link to="/shop" className="nav-item">View Items</Link>
					<Link to="/additem" className="nav-item">Add Item</Link>
				</div>
			</div>

			<Switch>
				<Route path="/login" component={Login} />
				<Route path="/shop" component={Shop} />
				<Route path="/additem" component={AddItem} />
			</Switch>
    </div>
  );
}

export default App;
