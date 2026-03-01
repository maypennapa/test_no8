import React from "react";
import { Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Button } from "@mui/material";

export default function ConfirmDeleteDialog({ 
  open,           
  onClose,       
  onConfirm,     
  title = "ยืนยันการลบ",           
  message = "คุณแน่ใจหรือไม่?",     
  confirmText = "ลบ",               
  cancelText = "ยกเลิก"             
}) {
  return (
    <Dialog open={open} onClose={onClose}>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>
        <DialogContentText>{message}</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>{cancelText}</Button>
        <Button color="error" onClick={onConfirm}>{confirmText}</Button>
      </DialogActions>
    </Dialog>
  );
}