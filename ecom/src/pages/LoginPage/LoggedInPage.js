import React from 'react';

const LoggedInPage = () => {
  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <h1>Вы залогинены, но не админ</h1>
    </div>
  );
};

export default LoggedInPage;