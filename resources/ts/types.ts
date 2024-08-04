export interface PageExpose {
  title: string
}

export enum UserRole {
  AdminRole = 'ADMIN',
  ManagerRole = 'MANAGER',
  SimpleUserRole = 'USER',
  GuestRole = 'GUEST',
}

export interface User {
  id: number
  email: string | null
  first_name: string | null
  last_name: string | null
  full_name: string | null
  avatar_id: number | null
  role: UserRole
  created_at: string
  updated_at: string
}

export interface Course {
  id: number
  uuid: string
  name: string
  user_id: number
  created_at: string
  updated_at: string
}