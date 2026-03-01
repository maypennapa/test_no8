import { render, screen } from '@testing-library/react'
import QuestionList from './QuestionList'
import { MemoryRouter } from 'react-router-dom'

test('แสดงหัวข้อ รายการข้อสอบ', () => {
  render(
    <MemoryRouter>
      <QuestionList />
    </MemoryRouter>
  )

  const heading = screen.getByText(/รายการข้อสอบ/i)
  expect(heading).toBeInTheDocument()
})