package grouphandler

import (
	"database/sql"
	"fmt"
	"splisense/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateGroup(group types.Group) (int, error) {
	result, err := s.db.Exec(
		"INSERT INTO groups (name, description, created_by) VALUES (?, ?, ?)",
		group.Name,
		group.Description,
		group.CreatedBy,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create group: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %v", err)
	}

	return int(id), nil
}

func (s *Store) AddGroupMember(member types.GroupMember) error {
	_, err := s.db.Exec(
		"INSERT INTO group_members (group_id, user_id) VALUES (?, ?)",
		member.GroupID,
		member.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to add group member: %v", err)
	}

	return nil
}

func (s *Store) GetGroupByID(id int) (*types.Group, error) {
	rows, err := s.db.Query(
		"SELECT id, name, description, created_by, created_at FROM groups WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query group: %v", err)
	}
	defer rows.Close()

	var group types.Group
	if rows.Next() {
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.Description,
			&group.CreatedBy,
			&group.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group: %v", err)
		}
	} else {
		return nil, fmt.Errorf("group not found")
	}

	return &group, nil
}

func (s *Store) GetGroupMembers(groupID int) ([]int, error) {
	rows, err := s.db.Query(
		"SELECT user_id FROM group_members WHERE group_id = ?",
		groupID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query group members: %v", err)
	}
	defer rows.Close()

	var members []int
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group member: %v", err)
		}
		members = append(members, userID)
	}

	return members, nil
}
