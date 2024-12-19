import daisyui from 'daisyui';

/** @type {import('tailwindcss').Config} */
export default {
  content: ['./intext.html', './src/**/*.{js,ts,jsx,tsx,vue}'],
  plugins: [daisyui],
  daisyui: {
    themes: [
      {
        budgyLight: {
          primary: '#FF6B35',
          secondary: '#1A9B7C',
          accent: '#4ECDC4',
          neutral: '#6C757D',

          'base-100': '#F8F9FA',
          'base-200': '#E9ECEF',
          'base-300': '#DEE2E6',
          'base-content': '#212529',

          'neutral-100': '#E9ECEF',
          'neutral-200': '#DEE2E6',
          'neutral-300': '#CED4DA',
          'neutral-400': '#ADB5BD',

          info: '#3498DB',
          success: '#2ECC71',
          warning: '#F39C12',
          error: '#E74C3C',

          'text-base': '#212529',
          'text-muted': '#6C757D',
          'text-disabled': '#ADB5BD',
        },
      },
      {
        budgyDark: {
          primary: '#FF6B35',
          secondary: '#1A9B7C',
          accent: '#4ECDC4',
          neutral: '#6C757D',

          'base-100': '#1A1E21',
          'base-200': '#212529',
          'base-300': '#343A40',
          'base-content': '#F8F9FA',

          'neutral-100': '#1A1E21',
          'neutral-200': '#212529',
          'neutral-300': '#343A40',
          'neutral-400': '#495057',

          info: '#3498DB',
          success: '#2ECC71',
          warning: '#F39C12',
          error: '#E74C3C',

          'text-base': '#F8F9FA',
          'text-muted': '#CED4DA',
          'text-disabled': '#6C757D',
        },
      },
    ],
  },
};
