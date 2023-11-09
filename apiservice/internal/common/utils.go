package common

//func (m *customCmsTopicCommentModel) Count(ctx context.Context) (int64, error) {
//	query := fmt.Sprintf("select count(*) as count from %s", m.table)
//
//	var count int64
//	err := m.conn.QueryRow(&count, query)
//
//	switch err {
//	case nil:
//		return count, nil
//	case sqlc.ErrNotFound:
//		return 0, ErrNotFound
//	default:
//		return 0, err
//	}
//}
