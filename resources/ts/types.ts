import { RouteLocationResolved } from 'vue-router'
import { Ref } from 'vue'

export interface PageExpose {
  title: string
  back?: RouteLocationResolved | Ref<RouteLocationResolved>
}

export enum UserRole {
  Guest = 'guest',
  Member = 'member',
  Company = 'company',
  Moderator = 'moderator',
  Admin = 'admin',
}

export const UserRoleTranslates: Record<UserRole, string> = {
  [UserRole.Guest]: 'Гость',
  [UserRole.Member]: 'Пользователь',
  [UserRole.Company]: 'Компания',
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

export const StatusBackgroundColors: Record<Status, string> = {
  [Status.Created]: 'bg-blue-300',
  [Status.Active]: 'bg-green-300',
  [Status.Inactive]: 'bg-red-300',
}

export const StatusTextColors: Record<Status, string> = {
  [Status.Created]: 'text-blue-700',
  [Status.Active]: 'text-green-700',
  [Status.Inactive]: 'text-red-700',
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
  moderation_reason: string | null
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
  options: Option[]
  type: QuestionType
  status: Status
  explanation: string | null
  moderation_reason: string | null
  course_id: number
  module_id: number | null
  created_at: string
  updated_at: string
  tags_ids: number[] | null
  // user_answers: UserAnswer[] | null
}

export interface Option {
  id: number
  content: string
  is_correct: boolean
}

export enum TestSessionType {
  SelectionSystem = 'selection-system',
  Marathon = 'marathon',
  Mistake = 'mistake',
  Module = 'module',
  Exam = 'exam',
}

export const TestSessionTypeTranslates: Record<TestSessionType, string> = {
  [TestSessionType.SelectionSystem]: 'Система подбора',
  [TestSessionType.Marathon]: 'Марафон',
  [TestSessionType.Mistake]: 'Ошибки',
  [TestSessionType.Module]: 'Модуль',
  [TestSessionType.Exam]: 'Экзамен',
}

export interface UserAnswer {
  id: number
  is_correct: boolean
  answer_data: {
    answer_id: number
    answers_ids: number[]
    plaintext: string
  }
  question_id: number
  test_session_id: number
  answered_at: string
}

export interface TestSession {
  id: number
  uuid: string
  name: string
  course_id: number
  course_uuid: string | null
  user_id: number
  last_question_id: number | null
  question_ids: number[]
  questions: Question[] | null
  type: TestSessionType
  correct: number
  incorrect: number
  deleted_at: string | null
  created_at: string
  updated_at: string
}

export interface Tag {
  id: number
  name: string
}
