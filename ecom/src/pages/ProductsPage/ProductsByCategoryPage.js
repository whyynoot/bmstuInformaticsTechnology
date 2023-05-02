import ProductCard from "../../components/ProductCard/ProductCard";
import React, { useState, useEffect } from 'react';
import { useParams } from "react-router-dom";

function ProductsByCategoriesPage() {
    const [products, setProducts] = useState([]);
    let name = useParams();
  
    useEffect(() => {
      async function fetchProducts() { 
        const response = await fetch(`/api/${name.name}/products`);
        const data = await response.json();
        if (data.status !== "OK" || data.products.length === 0) {
          setProducts([])
          return ;
        }
        setProducts(data.products)
      }
      fetchProducts();
    }, []);
    
      return (
          <section className="py-5">
              <div className="container px-4 px-lg-5 mt-5">
                  <div className="row gx-4 gx-lg-5 row-cols-2 row-cols-md-3 row-cols-xl-4 justify-content-center">
                      {products.map(product => (
                          ProductCard(product)
                      ))}
                  </div>
              </div>
          </section>
      );
  }

export default ProductsByCategoriesPage;