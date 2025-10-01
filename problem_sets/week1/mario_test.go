package week1

import "testing"

func TestMarioStair(t *testing.T) {
	t.Run("negative height", func(t *testing.T) {
		msg := MarioStair(-1)
		expected := ""
		if msg != expected {
			t.Errorf("MarioStair(-1) = %v, want empty string", msg)
		}
	})

	t.Run("height of 0", func(t *testing.T) {
		msg := MarioStair(0)
		expected := ""
		if msg != expected {
			t.Errorf("MarioStair(0) = %v, want empty string", msg)
		}
	})

	t.Run("height of 4", func(t *testing.T) {
		msg := MarioStair(4)
		expected := `   #
  ##
 ###
####
`
		if msg != expected {
			t.Errorf("MarioStair(4) = %v, want bricks as: %s", msg, expected)
		}
	})

	t.Run("height of 8", func(t *testing.T) {
		msg := MarioStair(8)
		expected := `       #
      ##
     ###
    ####
   #####
  ######
 #######
########
`
		if msg != expected {
			t.Errorf("MarioStair(8) = %v, want bricks as: %s", msg, expected)
		}
	})
}

func TestMarioStairWithGap(t *testing.T) {
	t.Run("negative height", func(t *testing.T) {
		msg := MarioStairsWithGap(-1)
		expected := ""
		if msg != expected {
			t.Errorf("MarioStairsWithGap(-1) = %v, want empty string", msg)
		}
	})

	t.Run("height of 0", func(t *testing.T) {
		msg := MarioStairsWithGap(0)
		expected := ""
		if msg != expected {
			t.Errorf("MarioStairsWithGap(0) = %v, want empty string", msg)
		}
	})

	t.Run("height of 4", func(t *testing.T) {
		msg := MarioStairsWithGap(4)
		expected := `   # #
  ## ##
 ### ###
#### ####
`
		if msg != expected {
			t.Errorf("MarioStairsWithGap(4) = %v, want bricks as: %s", msg, expected)
		}
	})

	t.Run("height of 8", func(t *testing.T) {
		msg := MarioStairsWithGap(8)
		expected := `       # #
      ## ##
     ### ###
    #### ####
   ##### #####
  ###### ######
 ####### #######
######## ########
`
		if msg != expected {
			t.Errorf("MarioStairsWithGap(8) = %v, want bricks as: %s", msg, expected)
		}
	})
}
