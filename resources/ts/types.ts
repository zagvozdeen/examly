export enum UserRole {
  AdminRole = 'ADMIN',
  ManagerRole = 'MANAGER',
  SimpleUserRole = 'USER',
  GuestRole = 'GUEST',
}

export interface User {
  id: number
  role: UserRole
  created_at: string
  updated_at: string
}

export interface Course {
  id: number
}