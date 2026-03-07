import { defineStore } from 'pinia';

interface User {
  email: string;
  role: string;
}

interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
}

export const useAuthStore = defineStore('auth', {
  persist: true,

  state: (): AuthState => ({
    isAuthenticated: false,
    user: null,
  }),

  actions: {
    login(user: User) {
      this.isAuthenticated = true;
      this.user = user;
    },

    logout() {
      this.isAuthenticated = false;
      this.user = null;
    },
  },
});
