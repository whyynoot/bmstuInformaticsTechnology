import ProductCard from "../../components/ProductCard/ProductCard";
import React, { useState, useEffect } from 'react';


function ProductsPage () {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    async function fetchProducts() {
      const response = await fetch('/api/products');
      const data = await response.json();
      if (data.status !== "OK") {
        setProducts([])
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

export default ProductsPage;

