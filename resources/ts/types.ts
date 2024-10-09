import { RouteLocationResolved } from 'vue-router'

export interface PageExpose {
  title: string
  back?: RouteLocationResolved
}

export enum UserRole {
  Guest = 'guest',
  Member = 'member',
  Moderator = 'moderator',
  Admin = 'admin',
}

export const UserRoleTranslates: Record<UserRole, string> = {
  [UserRole.Guest]: 'Гость',
  [UserRole.Member]: 'Пользователь',
  [UserRole.Moderator]: 'Модератор',
  [UserRole.Admin]: 'Администратор',
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

export enum Status {
  Created = 'created',
  Active = 'active',
  Inactive = 'inactive',
}

export const StatusTranslates: Record<Status, string> = {
  [Status.Created]: 'Новый',
  [Status.Active]: 'Активен',
  [Status.Inactive]: 'Неактивен',
}

export interface Course {
  id: number
  uuid: string
  name: string
  description: string | null
  moderation_reason: string | null
  color: string
  icon: string
  status: Status
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
  status: Status
  course_id: number
  user_id: number
  created_at: string
  updated_at: string
  course: Course | null
}

export enum QuestionType {
  SingleChoice = 'single_choice',
  MultipleChoice = 'multiple_choice',
  Plaintext = 'plaintext',
}

export const QuestionTypeTranslates: Record<QuestionType, string> = {
  [QuestionType.SingleChoice]: 'Один ответ',
  [QuestionType.MultipleChoice]: 'Несколько ответов',
  [QuestionType.Plaintext]: 'Поле для ввода',
}

export interface Question {
  id: number
  uuid: string
  title: string
  content: string | null
  type: QuestionType
  status: Status
  created_at: string
  updated_at: string
}

export enum TestSessionType {
  Marathon = 'marathon',
  Mistake = 'mistake',
  Module = 'module',
  Exam = 'exam',
}

export const TestSessionTypeTranslates: Record<TestSessionType, string> = {
  [TestSessionType.Marathon]: 'Марафон',
  [TestSessionType.Mistake]: 'Модуль',
  [TestSessionType.Module]: 'Ошибки',
  [TestSessionType.Exam]: 'Экзамен',
}

export interface UserCourse {
  id: number
  uuid: string
  name: string
  type: TestSessionType
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

export interface FullCourseStats {
  id: number
  uuid: string
  type: TestSessionType
  created_at: number
  correct: number
  incorrect: number
  total: number
}