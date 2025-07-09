import React from 'react';
import CardActivationForm from '../components/CardActivationForm';

const ActivateCardPage: React.FC = () => {
  return (
    <div
      style={{
        minHeight: '100vh',
        background: '#f2e9e4',
        padding: '40px 0',
      }}
    >
      <h1
        style={{
          color: '#4a4e69',
          textAlign: 'center',
          marginBottom: 32,
          letterSpacing: 1,
        }}
      >
        Activate Card
      </h1>
      <div
        style={{
          maxWidth: 420,
          margin: '0 auto',
          background: 'linear-gradient(135deg, #4a4e69 30%, #9a8c98 90%)',
          borderRadius: 16,
          boxShadow: '0 4px 24px #4a4e6922',
          padding: 32,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center'
        }}
      >
        <CardActivationForm />
      </div>
    </div>
  );
};

export default ActivateCardPage; 