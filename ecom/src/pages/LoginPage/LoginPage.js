import React, { useState } from 'react';
import { Container, Form, Button } from 'react-bootstrap';
import Cookies from 'js-cookie';
import LoggedInPage from './LoggedInPage';
import jwt_decode from "jwt-decode";
import AdminPage from '../AdminPage/AdminPage';
import { useRef } from 'react';

const LoginPage = () => {
    const usernameRef = useRef('');
    const passwordRef = useRef('');
    let isLoggedIn = false;
    let isAdmin = false;
    const token = Cookies.get('token');

    const handleSubmit = async (event) => {
        event.preventDefault();
        const username = usernameRef.current.value;
        const password = passwordRef.current.value;

        try {
          let response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              username,
              password,
            }),
          });
    
          const data = await response.json();
    
          if (response.ok) {
            console.log(response)
            Cookies.set('token', data.token);
            Cookies.set('refresh_token', data.refresh_token);
            window.location.reload();
          } else {
            console.log(data.message);
          }
        } catch (error) {
          console.log(error);
        }
      };

    if (token) {
        console.log(token);
        const decodedToken = jwt_decode(token, 'n5VFbBxdSjGiymiKaZvGeKzzOlMRR8G6');
        isAdmin = decodedToken.admin;
        if (isAdmin) {
            return <AdminPage />;
        } else {
            return <LoggedInPage />;
        }
    } else { 
        return (
            <Container className="d-flex justify-content-center align-items-center vh-100">
              <Form onSubmit={handleSubmit} className="bg-light p-5 rounded">
                <h1 className="text-center mb-4">Вход</h1>
                <Form.Group controlId="formBasicEmail">
                  <Form.Label>Username</Form.Label>
                  <Form.Control type="username" placeholder="Введите username" ref={usernameRef} />
                </Form.Group>
                <Form.Group controlId="formBasicPassword" className='mt-4'>
                  <Form.Label>Пароль</Form.Label>
                  <Form.Control type="password" placeholder="Введите пароль" ref={passwordRef} />
                </Form.Group>
                <Button variant="primary" type="submit" className="w-100 mt-5">
                  Войти
                </Button>
              </Form>
            </Container>
          );
    }
};

export default LoginPage;