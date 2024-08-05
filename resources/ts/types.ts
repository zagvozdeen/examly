import { RouteLocationResolved } from 'vue-router'

export interface PageExpose {
  title: string
  back?: RouteLocationResolved
}

export enum UserRole {
  AdminRole = 'ADMIN',
  ManagerRole = 'MANAGER',
  SimpleUserRole = 'USER',
  GuestRole = 'GUEST',
}

export const UserRoleTranslates: Record<UserRole, string> = {
  [UserRole.AdminRole]: 'Администратор',
  [UserRole.ManagerRole]: 'Модератор',
  [UserRole.SimpleUserRole]: 'Пользователь',
  [UserRole.GuestRole]: 'Гость',
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

export interface FileModel {
  id: number
  uuid: string
  content: string
  size: number
  mime_type: string
  origin_name: string
  created_at: string
  updated_at: string
}