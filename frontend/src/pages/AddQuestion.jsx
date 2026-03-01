import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api/axios';
import { Box, TextField, Button, Stack, Typography } from '@mui/material';

export default function AddQuestion() {
  const nav = useNavigate();

  const [form, setForm] = useState({
    question: '',
    choices: Array(4).fill(''),
  });

  const [errors, setErrors] = useState({
    question: false,
    choices: Array(4).fill(false),
  });

  const handleChange = (value, index) => {
    const newChoices = [...form.choices];
    newChoices[index] = value;
    setForm(prev => ({ ...prev, choices: newChoices }));

    const newErrors = [...errors.choices];
    newErrors[index] = false;
    setErrors(prev => ({ ...prev, choices: newErrors }));
  };

  const submit = async () => {
    const questionError = !form.question.trim();
    const choicesError = form.choices.map(c => !c.trim());

    setErrors({ question: questionError, choices: choicesError });

    if (questionError || choicesError.some(e => e)) return;

    try {
      const payload = { question: form.question };
      form.choices.forEach((c, i) => (payload[`choice${i + 1}`] = c));

      await api.post('/questions', payload);
      nav('/');
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <Box sx={{ p: 4, maxWidth: 600, mx: 'auto' }}>
      <Typography variant="h4" mb={3}>เพิ่มข้อสอบ</Typography>
      <Stack spacing={2}>
        <TextField
          label="คำถาม"
          value={form.question}
          onChange={e => {
            setForm(prev => ({ ...prev, question: e.target.value }));
            setErrors(prev => ({ ...prev, question: false }));
          }}
          fullWidth
          error={errors.question}
          helperText={errors.question ? 'กรุณากรอกคำถาม' : ''}
        />

        {form.choices.map((choice, index) => (
          <TextField
            key={index}
            label={`คำตอบ ${index + 1}`}
            value={choice}
            onChange={e => handleChange(e.target.value, index)}
            fullWidth
            error={errors.choices[index]}
            helperText={errors.choices[index] ? `กรุณากรอกคำตอบ ${index + 1}` : ''}
          />
        ))}

        <Stack direction="row" spacing={2} mt={2}>
          <Button variant="contained" color="primary" onClick={submit}>บันทึก</Button>
          <Button variant="outlined" color="secondary" onClick={() => nav('/')}>ยกเลิก</Button>
        </Stack>
      </Stack>
    </Box>
  );
}