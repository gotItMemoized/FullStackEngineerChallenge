import React, { lazy, Suspense } from 'react';
import { HashRouter, Route, Redirect, Switch } from 'react-router-dom';
import { Login } from './pages';
import useLocalStorage from 'react-use-localstorage';
import './App.scss';

// Codesplit so we don't pull in pages that we don't need/have permissions to
const AdminHome = lazy(() => import('./pages/adminHome'));
const UserHome = lazy(() => import('./pages/userHome'));
const Users = lazy(() => import('./pages/users'));
const NewUser = lazy(() => import('./pages/newUser'));
const EditUser = lazy(() => import('./pages/editUser'));
const AdminPerformanceSelector = lazy(() => import('./pages/adminPRSelection'));
const AdminPerformanceManager = lazy(() => import('./pages/adminPRManager'));
const AdminPerformanceNew = lazy(() => import('./pages/adminPRNew'));
const AdminPerformanceEdit = lazy(() => import('./pages/adminPREdit'));

const App = () => {
  const [user, setUser] = useLocalStorage(
    'jwt_token',
    JSON.stringify({ isAdmin: false, loggedIn: false }),
  );
  let routes;

  const setStringedUser = u => setUser(JSON.stringify(u));
  const parsedUser = JSON.parse(user);

  // an improvement we can make here is to hit a `/user/me` endpoint
  // in order to validate our token is valid before we make a determination
  //  on if they're admin or not. our api endpoints will always do the validation on request though

  if (parsedUser && parsedUser.loggedIn === true) {
    if (parsedUser.isAdmin) {
      routes = (
        <div>
          <Route
            exact
            path="/"
            component={() => <AdminHome user={parsedUser} setUser={setStringedUser} />}
          />
          <Route
            exact
            path="/users"
            component={() => <Users user={parsedUser} setUser={setStringedUser} />}
          />
          <Route
            exact
            path="/users/new"
            component={() => <NewUser currentUser={parsedUser} setCurrentUser={setStringedUser} />}
          />
          <Route
            exact
            path="/users/:id/edit"
            component={props => (
              <EditUser currentUser={parsedUser} setCurrentUser={setStringedUser} {...props} />
            )}
          />
          <Route
            exact
            path="/performance"
            component={() => (
              <AdminPerformanceSelector currentUser={parsedUser} setCurrentUser={setStringedUser} />
            )}
          />
          <Route
            exact
            path="/performance-manager"
            component={() => (
              <AdminPerformanceManager currentUser={parsedUser} setCurrentUser={setStringedUser} />
            )}
          />
          <Route
            exact
            path="/performance-manager/new"
            component={() => (
              <AdminPerformanceNew currentUser={parsedUser} setCurrentUser={setStringedUser} />
            )}
          />
          <Route
            exact
            path="/performance-manager/:id/edit"
            component={props => (
              <AdminPerformanceEdit
                currentUser={parsedUser}
                setCurrentUser={setStringedUser}
                {...props}
              />
            )}
          />
          <Route path="/login" component={() => <Redirect to="/" />} />
        </div>
      );
    } else {
      routes = (
        <div>
          <Route
            exact
            path="/"
            component={() => <UserHome user={parsedUser} setUser={setStringedUser} />}
          />
          <Route path="/login" component={() => <Redirect to="/" />} />
        </div>
      );
    }
  } else {
    routes = (
      <div>
        <Switch>
          <Route
            exact
            path="/login"
            component={() => <Login user={parsedUser} setUser={setStringedUser} />}
          />
          <Redirect to="/login" />
        </Switch>
      </div>
    );
  }
  return (
    <div className="App">
      <HashRouter>
        <Suspense fallback={'loading'}>{routes}</Suspense>
      </HashRouter>
    </div>
  );
};

export default App;
