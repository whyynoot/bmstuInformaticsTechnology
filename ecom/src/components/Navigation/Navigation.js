import React, { useState, useEffect } from 'react';
import { Navbar, Nav, Dropdown, Container, Button } from 'react-bootstrap';

function Navigation() {
  const [categories, setCategories] = useState([]);

  useEffect(() => {
    async function fetchCategories() {
      const response = await fetch('/api/categories');
      const data = await response.json();
      if (data.status !== "OK") {
        setCategories([])
      }
      setCategories(data.categories);
    }

    fetchCategories();
  }, []);

  return (
    <Navbar expand="lg" bg="dark" variant="dark">
      <Container>
        <Navbar.Brand href="/">Ecom App</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav>
            <Nav.Link href="/products">Продукты</Nav.Link>
            <Dropdown as={Nav.Item}>
              <Dropdown.Toggle as={Nav.Link}>Категории</Dropdown.Toggle>
              <Dropdown.Menu>
                {categories.map(category => (
                  <Dropdown.Item href={`/category/${category.api_name}`}>{category.name_ru}</Dropdown.Item>
                ))}
              </Dropdown.Menu>
            </Dropdown>
          </Nav>
          <Nav className="ms-auto">
            <Button variant="dark"><Nav.Link href="/login">Log in</Nav.Link></Button>
          </Nav>
          </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default Navigation;