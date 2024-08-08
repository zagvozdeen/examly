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

export enum CourseStatus {
  New = 'NEW',
  Active = 'ACTIVE',
  Inactive = 'INACTIVE',
}

export const CourseStatusTranslates: Record<CourseStatus, string> = {
  [CourseStatus.New]: 'Новый',
  [CourseStatus.Active]: 'Активен',
  [CourseStatus.Inactive]: 'Неактивен',
}

export interface Course {
  id: number
  uuid: string
  name: string
  description: string | null
  color: string
  icon: string
  status: CourseStatus
  user_id: number
  created_at: string
  updated_at: string
}

export type CourseStats = Array<{completed: number, total: number, name: string}>

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

export interface Module {
  id: number
  uuid: string
  name: string
  status: CourseStatus
  course_id: number
  user_id: number
  created_at: string
  updated_at: string
  course: Course | null
}

export enum QuestionType {
  OneAnswerType = 'ONE_ANSWER',
  MultiplyAnswersType = 'MULTIPLY_ANSWERS',
  InputType = 'INPUT',
}

export const QuestionTypeTranslates: Record<QuestionType, string> = {
  [QuestionType.OneAnswerType]: 'Один ответ',
  [QuestionType.MultiplyAnswersType]: 'Несколько ответов',
  [QuestionType.InputType]: 'Поле для ввода',
}

export interface Question {
  id: number
  uuid: string
  content: string
  type: QuestionType
  status: CourseStatus
  created_at: string
  updated_at: string
}

export enum UserCourseType {
  MarathonUserCourseType = 'MARATHON',
  ModuleUserCourseType = 'MODULE',
  ExamUserCourseType = 'EXAM',
}

export interface UserCourse {
  id: number
  uuid: string
  name: string
  type: UserCourseType
  user_id: number
  course_id: number
  last_question_id: number | null
  created_at: string
  updated_at: string
  modules: UserModule[]
  questions: UserQuestion[]
}

export interface UserModule {
  id: number
  name: string
  course_id: number
  created_at: string
  updated_at: string
}

export interface UserQuestion {
  id: number
  uuid: string
  content: string
  explanation: string
  is_true: boolean | null
  course_id: number
  module_id: number | null
  question_id: number
  file_id: number
  type: QuestionType
  created_at: string
  updated_at: string
  answers: UserAnswer[]
}

export interface UserAnswer {
  id: number
  content: string
  question_id: number
  is_chosen: boolean
  is_true: boolean
  created_at: string
  updated_at: string
}