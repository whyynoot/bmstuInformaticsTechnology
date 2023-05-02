import { Container, Row, Col, Image, Button } from 'react-bootstrap';
import { useState, useEffect } from 'react';
import { useParams } from "react-router-dom";


function ProductPage (props) {
    const [products, setProducts] = useState({});
    let name = useParams(); 

    useEffect(() => {
      async function fetchProducts() {
        const response = await fetch('/api/products');
        const data = await response.json();
        if (data.status !== "OK") {
          setProducts({})
        }
        console.log(data, name)
        let product = data.products.find(item => item.api_name === name.name);
        console.log(product)
        setProducts(product)
      }

      fetchProducts();
    }, []);

    return (
    <Container className="d-flex align-items-center mt-5 pt-5">
          <Row>
            <Col md={5} className="d-flex justify-content-center">
              <Image src={products.image_path} fluid className="product-image" />
            </Col>
            <Col md={7}>
              <div className="d-flex flex-column h-100 justify-content-center">
                <h1 className="text-center">{products.name}</h1>
                <p className="text-center">
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod purus vel diam dignissim, eget cursus
                  ligula blandit. Donec aliquam blandit augue, eu accumsan magna feugiat ac.
                </p>
                <div className="text-center">
                  <p>
                    Цена: <strong>{products.price}</strong> руб.
                  </p>
                  <Button variant="primary">Add to Cart</Button>
                </div>
              </div>
            </Col>
          </Row>
        </Container>
        );
};

export default ProductPage;