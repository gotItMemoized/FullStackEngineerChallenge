import React, { lazy, Suspense } from 'react';
import { HashRouter, Route, Redirect, Switch } from 'react-router-dom';
import { Login, Loading } from './pages';
import Header from './containers/header';
import useLocalStorage from 'react-use-localstorage';
import '../node_modules/bulma/css/bulma.min.css';

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
const AdminPerformanceView = lazy(() => import('./pages/adminPRView'));
const PerformanceReviews = lazy(() => import('./pages/performanceReviews'));
const PerformanceReview = lazy(() => import('./pages/performanceReview'));
const NotFound = lazy(() => import('./pages/notFound'));

export const userDefault = { isAdmin: false, loggedIn: false };

const App = () => {
  const [user, setUser] = useLocalStorage('jwt_token', JSON.stringify(userDefault));
  let routes = [];

  const setStringedUser = u => setUser(JSON.stringify(u || userDefault));
  const parsedUser = JSON.parse(user);

  // an improvement we can make here is to hit a `/user/me` endpoint
  // in order to validate if our token/permissions is valid before we do any routing
  // Regardless, the api endpoints will always do the validation on request

  if (parsedUser && parsedUser.loggedIn === true) {
    // common header, match to all pages
    routes.push(
      <Route
        key="*"
        path="*"
        component={() => <Header user={parsedUser} setUser={setStringedUser} />}
      />,
    );
    let viewPages = [];
    if (parsedUser.isAdmin) {
      // Admin only routes
      viewPages.push(
        <Route exact key="/" path="/" component={AdminHome} />,
        <Route exact path="/users" key="/users" component={() => <Users user={parsedUser} />} />,
        <Route exact key="/users/new" path="/users/new" component={NewUser} />,
        <Route exact key="/users/:id/edit" path="/users/:id/edit" component={EditUser} />,
        <Route exact key="/performance" path="/performance" component={AdminPerformanceSelector} />,
        <Route
          exact
          key="/performance-manager"
          path="/performance-manager"
          component={AdminPerformanceManager}
        />,
        <Route
          exact
          key="/performance-manager/new"
          path="/performance-manager/new"
          component={AdminPerformanceNew}
        />,
        <Route
          exact
          key="/performance-manager/:id/edit"
          path="/performance-manager/:id/edit"
          component={AdminPerformanceEdit}
        />,
        <Route
          exact
          key="/performance-manager/:id/view"
          path="/performance-manager/:id/view"
          component={AdminPerformanceView}
        />,
      );
    } else {
      // regular user only routes
      viewPages.push(
        <Route exact key="/" path="/" component={UserHome} />,
        <Route
          key="/performance"
          path="/performance"
          component={() => <Redirect to="/performance-reviews" />}
        />,
      );
    }
    // common user/admin routes
    viewPages.push(
      <Route
        exact
        key="/performance-reviews"
        path="/performance-reviews"
        component={PerformanceReviews}
      />,
      <Route
        exact
        key="/performance-reviews/:id"
        path="/performance-reviews/:id"
        component={PerformanceReview}
      />,
      <Route key="/login" path="/login" component={() => <Redirect to="/" />} />,
    );

    routes.push(
      <Switch key="pageOr404">
        {viewPages}
        <Route key="404" path="*" exact={true} component={NotFound} />,
      </Switch>,
    );
  } else {
    // route if you're not logged in :)
    routes.push(
      <Switch key="/login">
        <Route
          exact
          path="/login"
          component={() => <Login user={parsedUser} setUser={setStringedUser} />}
        />
        <Redirect to="/login" />
      </Switch>,
    );
  }
  return (
    <div className="App">
      <HashRouter>
        <Suspense fallback={<Loading user={parsedUser} setUser={setStringedUser} />}>
          {routes}
        </Suspense>
      </HashRouter>
    </div>
  );
};

export default App;
