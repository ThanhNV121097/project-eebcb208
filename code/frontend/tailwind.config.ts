import type { Config } from 'tailwindcss';

const config: Config = {
  content: ['./app/**/*.{ts,tsx}', './components/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        bg: 'var(--color-bg)',
        text: 'var(--color-text)',
        muted: 'var(--color-text-muted)',
      },
      spacing: {
        2: 'var(--space-2)',
        4: 'var(--space-4)',
      },
      fontFamily: {
        sans: 'var(--font-body)',
      },
    },
  },
  plugins: [],
};

export default config;
