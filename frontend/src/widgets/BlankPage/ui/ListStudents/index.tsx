import {
  Stack,
  List,
  ListItem,
  IconButton,
  Typography,
  Pagination,
  ListItemButton,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import { useAppDispatch } from "@/shared/lib/hooks/redux.ts";
import { useEffect, useMemo, useState } from "react";
import styles from "./index.module.css";
import { handleDelete } from "@/widgets/BlankPage/ui/ListStudents/lib/handles/deleteStudent.ts";
import { handleGetStudent } from "@/widgets/BlankPage/ui/ListStudents/lib/handles/getStudent.ts";
import { useGetStudents } from "@/widgets/BlankPage/ui/ListStudents/lib/hooks/useGetStudents.ts";

const PER_PAGE = 7;

export function ListStudents() {
  const dispatch = useAppDispatch();

  const [students] = useGetStudents();

  useEffect(() => {
    console.log(students);
  }, [students]);

  const [page, setPage] = useState(1);

  useEffect(() => setPage(1), [students]);

  const paginated = useMemo(() => {
    const start = (page - 1) * PER_PAGE;
    return students.slice(start, start + PER_PAGE);
  }, [students, page]);

  return (
    <Stack spacing={1} className={styles.container}>
      <h2 className="font-bold text-lg">История бланков</h2>
      <List className={styles.list}>
        {paginated.map((s) => (
          <ListItem
            key={s.id}
            disablePadding
            secondaryAction={
              <IconButton
                edge="end"
                onClick={() => handleDelete({ id: s.id, dispatch })}
              >
                <DeleteIcon />
              </IconButton>
            }
          >
            <ListItemButton
              className={styles.itemButton}
              onClick={() => handleGetStudent(s.id)}
            >
              <Typography className={styles.itemText}>{s.fullName}</Typography>
            </ListItemButton>
          </ListItem>
        ))}
      </List>

      <Pagination
        size="small"
        className={styles.pagination}
        page={page}
        count={Math.ceil(students.length / PER_PAGE)}
        onChange={(_, p) => setPage(p)}
        sx={{ alignSelf: "center" }}
      />
    </Stack>
  );
}
