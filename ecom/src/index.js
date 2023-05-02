import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider, 
  Routes
} from 'react-router-dom'
import Header from './components/Header/Header';
import 'bootstrap/dist/css/bootstrap.min.css';
import ProductsPage from './pages/ProductsPage/ProductsPage';
import ProductPage from './pages/ProductPage/ProductPage';
import ProductsByCategoriesPage from './pages/ProductsPage/ProductsByCategoryPage'
import LoginPage from './pages/LoginPage/LoginPage';

const router = createBrowserRouter(
  createRoutesFromElements(
      <Route>
        <Route index={true} element={<Header />} />
        <Route path="products" element={<ProductsPage/>} />
        <Route path="product/:name" element={<ProductPage/>} />
        <Route path="category/:name" element={<ProductsByCategoriesPage/>} />
        <Route path="login" element={<LoginPage/>} />
      </Route>
  )
);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <App></App>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
