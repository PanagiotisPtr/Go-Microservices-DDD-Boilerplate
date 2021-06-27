import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Details from "./pages/Details";

const App = () => {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Login</Link>
            </li>
            <li>
              <Link to="/register">Register</Link>
            </li>
            <li>
              <Link to="/details">Details</Link>
            </li>
          </ul>
        </nav>

        <Switch>
          <Route path="/" exact={true}>
            <Login />
          </Route>
          <Route path="/register">
            <Register />
          </Route>
          <Route path="/details">
            <Details />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
