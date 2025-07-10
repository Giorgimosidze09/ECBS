import React from 'react';
import CardActivationForm from '../components/CardActivationForm';

const ActivateCardPage: React.FC = () => {
  return (
    <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <h1
        style={{
          color: 'var(--color-primary)',
          textAlign: 'center',
          marginBottom: 32,
          letterSpacing: 1,
          fontWeight: 800,
          fontSize: 32,
        }}
      >
        Activate Card
      </h1>
      <div
        className="fade-in scale-hover"
        style={{
          maxWidth: 420,
          width: '100%',
          margin: '0 auto',
          background: 'linear-gradient(135deg, var(--color-primary) 30%, var(--color-accent) 90%)',
          borderRadius: 'var(--radius)',
          boxShadow: '0 4px 24px var(--color-shadow)',
          padding: 32,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          transition: 'box-shadow var(--transition)',
        }}
      >
        <CardActivationForm />
      </div>
    </div>
  );
};

export default ActivateCardPage; 