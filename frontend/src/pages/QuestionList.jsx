import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import api from '../api/axios';
import {
    Box, Button, Typography, Card, CardContent, Stack,
    RadioGroup, FormControlLabel, Radio
} from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import DeleteIcon from '@mui/icons-material/Delete';
import ConfirmDeleteDialog from "../component/ConfirmDeleteDialog";


export default function QuestionList() {
    const [questions, setQuestions] = useState([]);
    const [selected, setSelected] = useState({});

    const [openDialog, setOpenDialog] = useState(false);
    const [deleteId, setDeleteId] = useState(null);
    const [deleteQuestionText, setDeleteQuestionText] = useState("");

    const fetchData = async () => {
        try {
            const res = await api.get('/questions');
            setQuestions(Array.isArray(res.data) ? res.data : []);
        } catch (err) {
            console.error(err);
        }
    };

    useEffect(() => {
        fetchData();
    }, []);

    const deleteQ = async (id) => {
        try {
            await api.delete(`/questions/${id}`);
            fetchData();
        } catch (err) {
            console.error(err);
        }
    };

    const handleSelect = (qId, choice) => {
        setSelected(prev => ({ ...prev, [qId]: choice }));
    }

    const handleOpenDelete = (id, questionText) => {
        setDeleteId(id);
        setDeleteQuestionText(questionText);
        setOpenDialog(true);
    }

    const handleConfirmDelete = async () => {
        try {
            if (deleteId !== null) {
                await api.delete(`/questions/${deleteId}`);
                setOpenDialog(false);
                setDeleteId(null);
                setDeleteQuestionText("");
                fetchData();
            }
        } catch (err) {
            console.error(err);
        }
    }

    return (
        <Box sx={{ p: 4 }}>
            <Stack direction="column" spacing={4} mb={3}>
                <Typography variant="h4">รายการข้อสอบ</Typography>
                <Button
                    variant="contained"
                    color="primary"
                    component={Link}
                    to="/add"
                    size="large"
                    startIcon={<AddIcon />}
                    sx={{ px: 2, py: 0.5, width: '9rem', fontSize: '16px' }}
                >
                    เพิ่มข้อสอบ
                </Button>
            </Stack>

            <Stack spacing={2}>
                {Array.isArray(questions) && questions.map((q) => (
                    <Card key={q.id} variant="outlined">
                        <CardContent>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 1 }}>
                                <Typography variant="h6">{q.number}. {q.question}</Typography>
                                <Button
                                    variant="outlined"
                                    color="error"
                                    startIcon={<DeleteIcon />}
                                    onClick={() => handleOpenDelete(q.id, q.question)}
                                >
                                    ลบ
                                </Button>
                            </Box>

                            <RadioGroup
                                value={selected[q.id] || ''}
                                onChange={(e) => handleSelect(q.id, e.target.value)}
                            >
                                <FormControlLabel value={q.choice1} control={<Radio />} label={`${q.choice1}`} />
                                <FormControlLabel value={q.choice2} control={<Radio />} label={`${q.choice2}`} />
                                <FormControlLabel value={q.choice3} control={<Radio />} label={`${q.choice3}`} />
                                <FormControlLabel value={q.choice4} control={<Radio />} label={`${q.choice4}`} />
                            </RadioGroup>
                        </CardContent>
                    </Card>
                ))}
            </Stack>
            <ConfirmDeleteDialog
                open={openDialog}
                onClose={() => setOpenDialog(false)}
                onConfirm={handleConfirmDelete}
                title="ยืนยันการลบคำถาม"
                message={`คุณแน่ใจว่าจะลบคำถาม "${deleteQuestionText}" หรือไม่? ลบแล้วไม่สามารถกู้คืนได้`}
                confirmText="ลบ"
                cancelText="ยกเลิก"
            />
        </Box>
    );
}